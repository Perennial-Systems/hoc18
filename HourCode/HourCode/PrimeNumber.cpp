#include "stdafx.h"
#include "PrimeNumber.h"
#include <iostream>
#include <conio.h>
#include <string>
#include <stdio.h>

using namespace std;



bool isPrime(int number, int maxNumber, int interation)
{
    if (number == 0)
    {
        return false;
    }
    if (number == 1 || number == 2 || number == 3 || number == 5 || number == 7)
    {
        return true;
    }
    if (maxNumber == interation)
    {
        return true;
    }
    if (number % interation == 0)
    {
        return false;
    }
    return isPrime(number, maxNumber, interation + 1);
}

void getPrimeNymber(int startNumber, int endNumber)
{

    if (isPrime(startNumber, startNumber / 2 + 1, 2))
    {
        std::cout << endl << "Prime " << startNumber;
    }
    if (startNumber == endNumber)
    {
        return;
    }
    getPrimeNymber(startNumber + 1, endNumber);
}

int main1()
{
    int x = 0;
    int y = 100;
    getPrimeNymber(x, y);
    char a;
    cin >> a;
    return 0;
}
