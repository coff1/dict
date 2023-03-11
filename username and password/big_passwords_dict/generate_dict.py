import argparse
import os

parser = argparse.ArgumentParser()
parser.add_argument("--list", "-l", help="Username list file", required=False)
parser.add_argument("--user", "-u", help="Username", required=False, default="")
parser.add_argument("--input", "-i", help="Input dictionary file", required=False, default="big_password_dict.txt")
parser.add_argument("--output", "-o", help="Output dictionary file", required=False, default="big_password_dict_generate.txt")
args = parser.parse_args()

with open(args.input, encoding="utf-8") as f:
    passwords_input = {i.strip() for i in f}

if args.list:
    with open(args.list, encoding="utf-8") as f:
        usernames = {i.strip() for i in f}
else:
    usernames = {args.user}

password_output = set()
n = 0
for i in usernames:
    for j in passwords_input:
        if "%user" in j:
            j = j.replace("%user", i)
        if j not in password_output:
            password_output.add(j)
        n += 1
        print(f"{n}\r", end="")

with open(args.output, "w", encoding="utf-8") as f:
    f.write("\n".join(password_output))
