function prime(x, y) {
    if (y < x) {
        console.log('y should be greater than x')
        return;
    }
    if (!y || !x) {
        console.log('any of should not be zero')
        return;
    }

    if (!y || !x) {
        console.log('any of should not be zero')
        return;
    }

    if (x == 1) x++;

    interate(x, y);
}

function interate(x, y) {    

    if (x == y) { return; }
    if (x % 2 != 0 && x % 3 != 0 && x % 5 != 0 && x % 7 != 0) {
        console.log(x);        
    }   
    ++x;    
    interate(x, y)
}