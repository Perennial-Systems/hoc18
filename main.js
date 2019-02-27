class Encoding {
    constructor() {
        this.strArrray = []
        this.count = 0;
    }

    // Method to store the ASCII values of each character of the input string into an Array
    textToAscii(str) {
        for (let index = 0; index < str.length; index++) {
            this.strArrray.push(str[index].charCodeAt(0)); //why 0 ??
        }

        return this.asciiToHex(this.strArrray);
    }

    //Methos to cnvert the ASCII codes into Hexadecimal values
    asciiToHex(arr) {
        this.strArrray = [];
        for (let char of arr) {
            let hexCode = char.toString(16);
            this.strArrray.push(hexCode);

        }
        return this.hexToBinary(this.strArrray);
    }

    //Method to convert the Hexadecimal values into Binary format(8-bits)
    hexToBinary(arr) {
        this.strArrray = [];
        for (let hex of arr) {
            this.strArrray.push(parseInt(hex, 16).toString(2).padStart(8, '0'));
        }

        let binary = '';
        for (let bin of this.strArrray) {
            binary += bin;
        }
        return this.binaryCheck(binary);
    }

    //Method to make sextets out of octets stored as a binary string
    binaryCheck(binary) {
        let byte = 6;
        let binArray = [];

        for (let index = 0; index < binary.length; index += byte) {
            let bits = '';
            for (let num = index; num <= index + (byte-1); num++) {
                bits += binary.charAt(num);
            }

            binArray.push(bits);
        }

        return this.binaryToDec(binArray);
    }

    //Method to convert 6-bit binary sets into decimal
    binaryToDec(arr) {
        this.count = 0;
        let rem = arr[arr.length - 1].length % 3;
        this.strArrray = [];

        //In case the least significant bits are not the multiples of 3, so hey are made explicitly by 
        //concatenating "00" at the end as per requirement
        if (rem != 0) {
            for (let index = 0; index < rem; index++) {
                arr[arr.length - 1] += "00";
                this.count++;
            }
        }

        //eventually the 6-bit binary sets are then converted into decimal format
        for (let element of arr) {
            this.strArrray.push(parseInt(element, 2).toString(10));
        }

        return this.decimalToBase(this.strArrray);
    }

    //Method to get base64_String characters based on the decimal values
    decimalToBase(arr) {
        //a demo string is taken that represents the typical base64_String Table characters
        let baseString = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
        let code = "";
        let baseStringLength = 64;
        let padding = '';

        for (let index = 0; index < arr.length; index++) {
            if (arr[index] > (baseStringLength-1)) {
                code += baseString.charAt(arr[index % baseStringLength]); //incase the decimal value exceeds 63
            } else {
                code += baseString.charAt(arr[index]);
            }
        }

        //count is used to figure out how many "=" characters are needed to be padded at the end of code
        if (this.count == 1) {
            padding = "=";
            code += padding;
        } else if (this.count == 2) {
            padding = "=="
            code += padding;
        }
        return code;
    }
}

/* let strEncode = new Encoding(); */