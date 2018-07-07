<?php

// full processing procedure: 1:24.53
// only copy each cell as is: 1:07.78
// only reverse columns order: 1:07.64
// only reverse string in each cell: 1:09.33
// only uppercase each cell: 1:12.25
// only removing all digits from values using regexp: 1:18.01

$fdIn = \fopen('in.csv', 'r');

if (!$fdIn) {
    throw new \Exception('Can\'t open file - in.csv.');
}

$fdOut = \fopen('out.csv', 'w');

if (!$fdOut) {
    throw new \Exception('Can\'t open file - out.csv.');
}

while (($rowIn = \fgetcsv($fdIn)) !== false) {
    $rowOut = [];

    foreach ($rowIn as $valueIn) {
        $valueOut = '(NULL VALUE)';

        if ($valueIn !== null && $valueIn !== '') {
            $valueOut = \strrev($valueIn);
            $valueOut = \preg_replace('/\d/', '', $valueOut);
            $valueOut = \strtoupper($valueOut);
        }

        $rowOut[] = $valueOut;
    }

    \fputcsv($fdOut, array_reverse($rowOut));
}

//\fflush($fdOut);
//\ftruncate($file, \ftell($file));

\fclose($fdIn);
\fclose($fdOut);

exit(0);
