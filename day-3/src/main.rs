use std::collections::HashSet;


fn pt1() {
    let mut sum = 0;
    include_str!("./input.txt").lines().for_each(|line|
    {
        let cmpt_size = line.chars().count() / 2;

        let (cmpt1, cmpt2) = line.split_at(cmpt_size);
        let first_set: HashSet<char> = cmpt1.chars().collect();
        let second_set: HashSet<char> = cmpt2.chars().collect();
        let common = first_set.intersection(&second_set).nth(0).unwrap();

        let value = match common {
            'a'..='z' => *common as u32 - 'a' as u32,
            'A'..='Z' => *common as u32 - 'A' as u32 + 26,
            _ => panic!("{common}"),
        };

        sum += value + 1;
    });

    println!("{}", sum);
}

fn pt2() {
    let mut sacks = include_str!("./input.txt").lines();
    let mut sum = 0;
    
    loop {
        let line1 = sacks.next().unwrap_or("_");

        if line1 == "_" {
            break;
        }

        let line2 = sacks.next().unwrap();
        let line3 = sacks.next().unwrap();

        let first_set: HashSet<char> = line1.chars().collect();
        let second_set: HashSet<char> = line2.chars().collect();
        let third_set: HashSet<char> = line3.chars().collect();

        let common_1: HashSet<char>  = first_set.intersection(&second_set).cloned().collect();
        let common = common_1.intersection(&third_set).nth(0).unwrap();

        let value = match common {
            'a'..='z' => *common as u32 - 'a' as u32,
            'A'..='Z' => *common as u32 - 'A' as u32 + 26,
            _ => panic!("{common}"),
        };

        sum += value + 1;
    }

    println!("{}", sum);
}

fn main() {
    pt1();
    pt2();
}
