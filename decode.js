class Decoding {
    constructor() {
        this.dummyArray = [];
    }

    //Method to store the base64 equivalent decimal values of the encoded characters into an empty array
    baseToDecimal(str) {
        console.log(str);
        let base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
        let dummyArray = [];

        if (str[str.length - 2] === "=") {  //incase there are 2 "=" characters in the padding
            for (let i = 0; i < (str.length - 2); i++) {

                this.dummyArray.push(base64.indexOf(str[i]));

            }
        } else if (str[str.length - 1] === "=") {   //incase there is a single "=" character in the padding
            for (let i = 0; i < (str.length - 1); i++) {

                this.dummyArray.push(base64.indexOf(str[i]));

            }
        } else {
            for (let i = 0; i < (str.length); i++) {  //incase there is a none "=" character in the padding

                this.dummyArray.push(base64.indexOf(str[i]));

            }
        }
        
        return this.decimalToBinary(this.dummyArray);
    }

    //Method to convert the decimal values into Binary(6-bit) format
    //and prepend the required "0's" in order to make the 6-bit sets
    decimalToBinary(arr) {
        this.dummyArray = [];
        let binary = '';

        for (let i = 0; i < arr.length; i++) {
            this.dummyArray.push(arr[i].toString(2));

            for (let j = this.dummyArray[i].length; j != 6 && j < 7; j++) {
                this.dummyArray[i] = "0" + this.dummyArray[i];
            }

            binary += this.dummyArray[i];
        }

        return this.binarySet(binary);
    }


    //6-Bit Binary Sets are eventually converted into 8-bit Binary sets
    binarySet(binary) {
        this.dummyArray = [];
        for (let i = 0; i < binary.length; i += 8) {
            let bits = '';

            for (let j = i; j <= i + 7; j++) {
                bits += binary.charAt(j);
            }
            
            this.dummyArray.push(bits);
        }
        
        //last element of the array is checked to be a "00" or "0000" and is popped from the array
        if (this.dummyArray[this.dummyArray.length - 1].length !== 8) {
            this.dummyArray.pop(this.dummyArray[this.dummyArray.length - 1]);
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
        console.log(decodedText);
        return decodedText;
    }
}