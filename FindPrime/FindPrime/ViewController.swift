//
//  ViewController.swift
//  SR_Pushpak
//
//  Created by Perennial on 05/12/18.
//  Copyright © 2018 Perennial. All rights reserved.
//

import UIKit

class ViewController: UIViewController {
    
    var primeNumbers: [Int]!
    var number = 0
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
        
        
        getPrimeNumber(between: 1, y: 11, base: 2)
    }
    
    func getPrimeNumber(between x: Int, y: Int, base: Int) {
        var constBase = 2
        number = x
        if x > y {
            print("Limit reached")
        }
        
        if base >= y {
            print("\(x)\n")
            number += 1
            getPrimeNumber(between: number, y: y, base: constBase)
        } else if (x<=y) {
            if ((x % base == 0) && (x != base)) {
                number += 1
                getPrimeNumber(between: number, y: y, base: constBase)
            } else {
                constBase = base + 1
                getPrimeNumber(between: number, y: y, base: constBase)
            }
        }
        
    }
}
