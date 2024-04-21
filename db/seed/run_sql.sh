#!/bin/bash

SQL_FILE="test.sql"

make -f ../../makefile db-cli < "$SQL_FILE"
