use std::fs::File;
use std::io;
use std::io::prelude::*;

fn calculate(filename: &str) -> io::Result<Vec<i32>> {
    let file = File::open(&filename)?;
    let reader = io::BufReader::new(file);

    let mut cal_counts = Vec::new();
    let mut cal_count = 0;

    for line in reader.lines() {
        let line = line?;

        if line.is_empty() {
            cal_counts.push(cal_count);
            cal_count = 0;
        }
        else {
            cal_count += line.parse::<i32>().unwrap();
        }
    }
    Ok(cal_counts)
}

fn main() {
    if let Ok(mut _c) = calculate("src/input.txt") {
        _c.sort_unstable();
        println!("{}", _c.into_iter().rev().take(3).sum::<i32>());
    }


    // method #2 
    // split by new empty lines
    // for each set of cals loop through them and add their cals
    // collect them to a vec and sort


    // let mut cals = include_str!("../input.txt")
    // .split("\n\n")
    // .map(|e| e.lines().map(|c| c.parse::<u32>().unwrap()).sum())
    // .collect::<Vec<u32>>();
    // cals.sort_unstable();
    // println!("{}", cals.into_iter().rev().take(3).sum::<u32>());
}


