#!/bin/sh
#
# Solution for Project Euler problem 1, written in POSIX shell

sum=0
i=1; while [ $i -le 999 ]; do
    if [ $(( i % 3 )) -eq 0 ]; then
        sum=$(( sum + i ))
    elif [ $(( i % 5 )) -eq 0 ]; then
        sum=$(( sum + i ))
    fi
    i=$(( i + 1 ))
done
echo "$sum"
