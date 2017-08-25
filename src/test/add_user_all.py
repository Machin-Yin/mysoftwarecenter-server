#!/usr/bin/env python3
# -*- coding: UTF-8 -*-

import subprocess

file = open("user")
while 1:
    lines = file.readlines()
    if not lines:
        break

    for line in lines:
        line = line.strip('\n     ')
        (p1, p2, p3) = (line.split('|'))
        if p3 == '':
            p3 = 'Other'
        subprocess.call(["./add_user.sh", p1, p2, p3])

