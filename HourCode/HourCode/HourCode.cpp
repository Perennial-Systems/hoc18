// HourCode.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"
#include "targetver.h"
#include <iostream>
#include <conio.h>
#include <string>
#include <stdio.h>

#include <iostream>
#include <sstream> 

using namespace std;



void decodeString(string str)
{
    string alphabet[] = {"","a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z" };
    
    string temp = "";
    int addition = 0;
    int found;
    for (int i = 0; i < str.length();i++)
    {
         temp = str[i];
        if (temp == " ") 
        {
            cout << alphabet[addition];
            addition = 0;
        }
        if (stringstream(temp) >> found)
        {
            addition = addition + found;
        }
        temp = "";
    }
    cout << alphabet[addition];
}


int main()
{
    string inputString = "aWk60#ase2> o_0^4alis5L";
    decodeString(inputString);
    char a;
    cin >> a;
    return 0;
}