#!/bin/sh
#
# Solution for Project Euler problem 2, written in POSIX shell

fib1=1
fib2=1
sum=0
i=1; while [ $i -lt 4000000 ]; do
    i=$(( fib1 + fib2 ))
    fib1=$fib2
    fib2=$i
    if [ $(( i % 2 )) -eq 0 ]; then
        sum=$(( sum + i ))
    fi
done
echo "$sum"
