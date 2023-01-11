use std::collections::HashMap;

fn main() {
    let score = HashMap::from([
        ("A X" , 4),
        ("A Y" , 8),
        ("A Z" , 3),
        ("B X" , 1),
        ("B Y" , 5),
        ("B Z" , 9),
        ("C X" , 7),
        ("C Y" , 2),
        ("C Z" , 6),
    ]);

    let input = include_str!("./input.txt");
    let mut ttl_score = 0;

    for line in input.lines() {
        ttl_score += score[line];
    }

    println!("{}", ttl_score);
}

