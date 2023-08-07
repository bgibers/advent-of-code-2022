pub fn jobs() -> &'static [(fn(), &'static str)] {
    &[
        (day1::main, "day-1"),
        (day2::main, "day-2"),
    ]
}