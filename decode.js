class Decoding {
    constructor() {
        this.dummyArray = [];
    }

    //Method to store the base64 equivalent decimal values of the encoded characters into an empty array
    baseToDecimal(str) {
        const base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
        this.dummyArray = [];
        const stringLength = str.length;
        let secondLastCharacter = str[stringLength - 2];
        let lastCharacter = str[stringLength- 1];
        

        if (secondLastCharacter === "=") {  //incase there are 2 "=" characters in the padding
            for (let char = 0; char < (stringLength - 2); char++) {

                this.dummyArray.push(base64.indexOf(str[char]));

            }
        } else if (lastCharacter === "=") {   //incase there is a single "=" character in the padding
            for (let char = 0; char < (stringLength - 1); char++) {

                this.dummyArray.push(base64.indexOf(str[char]));

            }
        } else {
            for (let char = 0; char < (stringLength); char++) {  //incase there is no "=" character in the padding

                this.dummyArray.push(base64.indexOf(str[char]));

            }
        }
        
        return this.decimalToBinary(this.dummyArray);
    }

    //Method to convert the decimal values into Binary(6-bit) format
    //and prepend the required "0's" in order to make the 6-bit sets
    decimalToBinary(arr) {
        this.dummyArray = [];
        let binary = '';
        let byte = 6;

        for (let index = 0; index < arr.length; index++) {
            this.dummyArray.push(arr[index].toString(2));

            for (let decLength = this.dummyArray[index].length; decLength != byte && decLength < (byte+1); decLength++) {
                this.dummyArray[index] = "0" + this.dummyArray[index];
            }

            binary += this.dummyArray[index];
        }

        return this.binarySet(binary);
    }


    //6-Bit Binary Sets are eventually converted into 8-bit Binary sets
    binarySet(binary) {
        this.dummyArray = [];
        let byte = 8;
        let lastArrElement;
        for (let bit = 0; bit < binary.length; bit += byte) {
            let bits = '';

            for (let num = bit; num <= bit + (byte-1); num++) {
                bits += binary.charAt(num);
            }
            
            this.dummyArray.push(bits);
        }

        lastArrElement = this.dummyArray[this.dummyArray.length - 1];

        //last element of the array is checked to be a "00" or "0000" and is popped from the array
        if (lastArrElement.length !== byte) {
            this.dummyArray.pop(lastArrElement);
        }
        
        return this.binaryToText(this.dummyArray);
    }

    //the final 8-Bit Binary sets are then converted into Hexadecimal values
    //and then further converted to Text characters using their ASCII values
    binaryToText(arr) {
        this.dummyArray = [];
        let decodedText = '';
        for (let element of arr) {
            let hex = parseInt(element, 2).toString(16)
            decodedText += (String.fromCharCode(parseInt(hex, 16)));
        }
        return decodedText;
    }
}