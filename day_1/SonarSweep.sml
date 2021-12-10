(* AoC Day 1: Sonar Sweep in sml-nj *)

(* val measurements = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]; *)
use "Measurements.sml";

(* Part One *)

(* Counts how many measurements are greater than the previous. *)
fun countGreater (xs: int list) = countHelper 0 NONE xs
and countHelper n _ nil = n
|   countHelper n NONE (x::xs) = countHelper n (SOME x) xs
|   countHelper n (SOME last) (x::xs) =
        if x > last then countHelper (n + 1) (SOME x) xs
        else countHelper n (SOME x) xs;

val result = countGreater measurements;
print ((Int.toString result) ^ " results where greater than the previous!\n");

(* Part Two *)

(* Splits a list into a list of lists containing each three elements of the list. *)
fun split3 (x1::x2::x3::xs) = [x1, x2, x3]::split3 (x2::x3::xs)
|   split3 _ = nil;

(* Sums up a list of integers. *)
val sum = foldr (fn (x, y) => x + y) 0;

(* Counts how many sums of measurements in the sliding window increased. *)
fun sumGreater xs = countGreater (map sum (split3 xs));

val result = sumGreater measurements;
print ((Int.toString result) ^ " sums where greater than the previous!\n");
