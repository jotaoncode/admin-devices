#!/usr/bin/env bash
if [ $# -eq 0 ]
  then
    echo "No arguments supplied"
    exit 1
fi
TIME=$(date +%s)
touch ./sql/${TIME}_$1.up.sql
touch ./sql/${TIME}_$1.down.sql