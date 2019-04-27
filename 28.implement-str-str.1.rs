/*
 * @lc app=leetcode id=28 lang=rust
 *
 * [28] Implement strStr()
 */
impl Solution {
    pub fn str_str(haystack: String, needle: String) -> i32 {
        if haystack.len() < needle.len() {
            return -1
        }

        match &haystack[..needle.len()].to_string() {
            x if x == &needle => 0,
            _ => {
                let i = Solution::str_str(haystack[1..].to_string(), needle);

                if (i == -1) { -1 } else { i + 1 }
            }
        }
    }
}
