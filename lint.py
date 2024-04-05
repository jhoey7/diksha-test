'''
this module runs pylint on all python scripts found in a directory tree
'''

import asyncio
import os
import platform
import re
import subprocess
import sys
from concurrent.futures.thread import ThreadPoolExecutor

count = 0

async def check_async(output_file, module, semaphore):
    '''apply pylint to the file specified if it is a *.py file'''
    global count
    if module[-3:] == ".go":

        async with semaphore:
            print(f"CHECKING {module}")
            command = f"golint {module}"
            process = await asyncio.create_subprocess_shell(command, stdout=asyncio.subprocess.PIPE,
                                                            stderr=asyncio.subprocess.PIPE)
            stdout, stderr = await process.communicate()
            data = stdout.decode('utf-8')

            for line in data.splitlines():
                if line and "don't use ALL_CAPS" not in line:
                    # Exclude ALL_CAPS restriction from counter
                    count += 1

            return data


async def process_files_async(base_directory, output_file, semaphore):
    coroutines = []
    for root, dirs, files in os.walk(base_directory):
        for name in files:
            filepath = os.path.join(root, name)
            if 'vendor' not in filepath:
                coroutines.append(check_async(output_file, filepath, semaphore))

    completed, pending = await asyncio.wait(coroutines, return_when=asyncio.ALL_COMPLETED)

    with open(output_file, 'a', encoding='utf-8') as infile:
        for item in completed:
            result = item.result()
            if result:
                infile.write(result)
                print(result)


if __name__ == "__main__":
    BASE_DIRECTORY = os.getcwd()

    try:
        print(sys.argv)
        OUTPUT_FILE = sys.argv[1]
    except IndexError:
        OUTPUT_FILE = 'pylint.log'

    print("looking for *.go scripts in subdirectories of ", BASE_DIRECTORY)
    # for root, dirs, files in os.walk(BASE_DIRECTORY):
    #     for name in files:
    #         filepath = os.path.join(root, name)
    #         check(OUTPUT_FILE, filepath)

    if platform.system() == 'Windows':
        loop = asyncio.ProactorEventLoop()
        asyncio.set_event_loop(loop)
    else:
        loop = asyncio.get_event_loop()

    executor = ThreadPoolExecutor(max_workers=2)
    loop.set_default_executor(executor)

    semaphore = asyncio.Semaphore(4, loop=loop)
    loop.run_until_complete(process_files_async(BASE_DIRECTORY, OUTPUT_FILE, semaphore))

    loop.close()

    print("Done linting!")
    print("==" * 50)
    print("%d errors found" % count)
