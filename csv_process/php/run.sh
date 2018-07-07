#!/bin/bash
/usr/bin/time -v docker run -it --rm --name php-csv-processing -v "$PWD":/usr/src/myapp -w /usr/src/myapp php:7.2 php index.php
