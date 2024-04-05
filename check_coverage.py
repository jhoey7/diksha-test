import os
import subprocess
import sys
import hashlib

def check_lint():

    project_name = ('diksha-test')
    env_build = project_name.split('-')[0].upper()

    stdoutdata = subprocess.getoutput("python3 lint.py | grep 'errors found'")
    total_error = stdoutdata.split()[0]

    if total_error <= "0":
        print('success lint')
    else:
        print('failed lint')
        sys.exit("result linter : " + total_error + 'failed to deploy')

    job = subprocess.getoutput("go test diksha-test/... -v -cover | grep 'coverage' | grep -v 'ok' | awk '{print $2}' ")
    job_count = subprocess.getoutput("go test diksha-test/... -v -cover | grep 'coverage' | grep -v 'ok' | wc -l")
    job_run = subprocess.getoutput("go test diksha-test/... -v -cover | grep 'RUN' | wc -l")
    job_pass = subprocess.getoutput("go test diksha-test/... -v -cover | grep 'PASS:' | wc -l")

    job_count = parseResult(job_count)
    job_run = parseResult(job_run)
    job_pass = parseResult(job_pass)

    sum_prosess = int(job_count)
    out = 0
    result = ''
    test_coverage = 0
    for data in job.split('\n'):
        if data.startswith("go:") is False:
            result = (data.split(' ')[0].replace('%', ''))
            out = out + float(result)
            test_coverage = (out / (sum_prosess))
    result_message = str(result)
    unit_test_result = job_pass + '/' + job_run

    print(result_message)

    print("message summary")
    print(job_run + '/' + job_pass)

    if job_run == job_pass:
        print('success test project')
    else:
        print(unit_test_result)
        sys.exit("test coverage : " + test_coverage + result_message + 'unit test coverage :' + str(
            unit_test_result) + 'failed to deploy')

    if test_coverage >= 86:
        print('success result test')
        print(test_coverage)
    else:
        print(test_coverage)
        sys.exit("test coverage : " + str(test_coverage) + result_message  + 'failed to deploy')

def parseResult(output):
    out = ""
    for data in output.split('\n'):
        if data.startswith("go:") is False:
            out = data
        else:
            out = "0"
    return out

def run():
    check_lint()

if __name__ == "__main__":
    run()