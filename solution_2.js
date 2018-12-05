function decode (str){
    let intArray=Array.from(str).filter(a=>isNaN(parseInt(a,0)));
    return intArray.map(m=>String.fromCharCode(96+m)).join()

}