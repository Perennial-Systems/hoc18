<?php
checkPrimeInRange(1, 99);

function checkPrimeInRange($x, $y)
{
    if ($x > 1) {
        if (in_array($x, array(2, 3, 5, 7, 11))) {
            echo "\r\n";
            echo $x;
        }

        if ($x % 2 >= 1) {
            if ($x % 3 >= 1) {
                if ($x % 5 >= 1) {
                    if ($x % 7 >= 1) {
                        if ($x % 11 >= 1) {
                            echo "\r\n";
                            echo $x;
                        }
                    }
                }
            }
        }
    }

    if ($x == $y)
        return;

    checkPrimeInRange(++$x, $y);
}

?>