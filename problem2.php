<?php
$str = "adsaasds545 dsaddasda1231 23123";
$exploded_str = explode(" ", $str );
foreach ($exploded_str as $item){
    $sum = 0;
    $str_ary = str_split($item,1);
    foreach ($str_ary as $key => $val) {
        if(filter_var($val, FILTER_VALIDATE_INT)){
            $sum+= $val;
        }
    }
    var_dump($sum);
}

?>