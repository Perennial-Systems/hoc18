<?php
$str = "aWk60#ase2> o_2^293alis874L";
$exploded_str = explode(" ", $str );
foreach ($exploded_str as $item){
    $sum = 0;
    $str_ary = str_split($item,1);
    foreach ($str_ary as $key => $val) {
        if(filter_var($val, FILTER_VALIDATE_INT)){
            $sum+= $val;
        }
    }

    if($sum > 26){
        $sum = ($sum % 26);
    }

    echo chr(($sum-1)+65);
}

?>