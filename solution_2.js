const str = "aWk60#ase2> o_0^4alis5L"
        const myChars = str.split('')

        let numbers = myChars.filter( (char) => {
            const isNumber = isNaN(char) == false
            if (isNumber == true) {
                return String.fromCharCode(97 + n)
            }             
        })        
        
            console.log("String is: ", numbers.toString());