<?php

define('CSV_ROWS_COUNT', 2000000);
define('CSV_COLUMNS_COUNT', 20);
define('CSV_VALUE_MAX_LENGTH', 50);

$fd = \fopen('in.csv', 'w');

if (!$fd) {
    throw new \Exception('Can\'t create/open file.');

    exit(1);
}

for ($i = 0; $i < CSV_ROWS_COUNT; $i++) {
    $row = [];

    for ($j = 0; $j < CSV_COLUMNS_COUNT; $j++) {
        $valueLength = \mt_rand(0, CSV_VALUE_MAX_LENGTH);
        $value = null;
        if ($valueLength > 0) {
            $value = \generateRandomString($valueLength);
        }
        $row[] = $value;
    }

    \fputcsv($fd, $row);
}

\fclose($fd);

exit(0);

function generateRandomString($length) {
    return \substr(\str_shuffle(\str_repeat($x = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ', \ceil($length / \strlen($x)) )), 1, $length);
}

