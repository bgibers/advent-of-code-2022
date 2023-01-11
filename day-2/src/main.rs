use std::collections::HashMap;

fn main() {
    let score = HashMap::from([
        ("A X" , 0 + 3),
        ("A Y" , 3 + 1),
        ("A Z" , 6 + 2),
        ("B X" , 0 + 1),
        ("B Y" , 3 + 2),
        ("B Z" , 6 + 3),
        ("C X" , 0 + 2),
        ("C Y" , 3 + 3),
        ("C Z" , 6 + 1),
    ]);

    let input = include_str!("./input.txt");
    let mut ttl_score = 0;

    for line in input.lines() {
        ttl_score += score[line];
    }

    println!("{}", ttl_score);
}

