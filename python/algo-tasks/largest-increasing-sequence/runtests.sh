#!/bin/bash

PY_PRG=largest-inc-seq.py
for i in $(ls tests); do
    if [[ $i == *.in ]]; then
        OUT=$(cat tests/$i | python3 $PY_PRG)
        echo "TEST $i OUTPUT:" $OUT
        EXP_TESTFILE=${i%.*}.out
        echo "EXPECTED:" $(cat tests/$EXP_TESTFILE)
    fi
done
