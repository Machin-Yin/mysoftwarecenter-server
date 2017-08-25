#!/usr/bin/env python3
# -*- coding: UTF-8 -*-

import subprocess

file = open("banner")
while 1:
    lines = file.readlines()
    if not lines:
        break

    for line in lines:
        line = line.strip('\n     ')
        (p1, p2) = (line.split('|'))
        subprocess.call(["./add_banner.sh", p1, p2])

