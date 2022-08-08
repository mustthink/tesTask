#!/bin/bash
    set -e
    set -x

    echo "__________LOADING TABLES...__________"
    pg_restore -C -d dvdrental /var/lib/postgresql/dvdrental.tar