# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
    l = nums.length
    (0..(l - 2)).each do |i|
        ((i + 1)..(l - 1)).each do |j|
            if (nums[i] + nums[j]) == target
                return [i, j]
            end
        end
    end
end

puts two_sum([1,2,3,4,9, 20, 11, 8], 10)
