use std::collections::HashSet;


fn main() {
    let mut sum = 0;
    let sacks = include_str!("./input.txt").lines().for_each(|line|
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
