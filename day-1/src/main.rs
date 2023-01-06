use std::fs::File;
use std::io;
use std::io::prelude::*;

fn calculate(filename: &str) -> io::Result<i32> {
    let file = File::open(&filename)?;
    let reader = io::BufReader::new(file);

    let mut highest_cal_cnt = 0;
    let mut cal_count = 0;

    for line in reader.lines() {
        let line = line?;

        if line.is_empty() {
            if cal_count > highest_cal_cnt {
                highest_cal_cnt = cal_count;
            }
            cal_count = 0;
        }
        else {
            cal_count += line.parse::<i32>().unwrap();
        }
    }
    Ok(highest_cal_cnt)
}

fn main() {
    if let Ok(_c) = calculate("src/input.txt") {
        println!("{}", _c)
    }
}


