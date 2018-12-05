package com.hourofcode;

public class DecodeString {

	public static void main(String[] args) {

		String sampleStr = "aWk60#ase2> o_2^293alis874L";
		String[] stringToDecodeArr = sampleStr.split(" ");

		int[] decodedArray = new int[stringToDecodeArr.length];
		for (int i=0; i<stringToDecodeArr.length; i++) {
			int decodedDigit = getDecodedDigit(stringToDecodeArr[i]);
			decodedArray[i] = decodedDigit;
		}
		for (int i = 0; i < decodedArray.length; i++) {
			int decodedDigit = decodedArray[i];
			if(decodedDigit == 0) {
				System.out.print("");
			} else {
				System.out.print(Character.toString((char) (decodedArray[i] + 96)));
			}
		}
	}

	private static int getDecodedDigit(String sampleStr) {
		char[] charArray = sampleStr.toCharArray();
		int sum = 0;
		for (char c : charArray) {
			try {
				int intValue = Integer.parseInt(Character.toString(c));
				sum += intValue;
			} catch (Exception e) {

			}
		}
		if (sum > 26) {
			return sum % 26;
		} else {
			return sum;
		}
	}


}
