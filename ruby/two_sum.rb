# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
    l = nums.length
    h = {}
    (0..(l - 1)).each do |i|
        h[nums[i]] = i
    end

    (0..(l - 1)).each do |i|
        k = target - nums[i]
        if h.key?(k) and h[k] != i then
            return [i, h[k]].sort
        end
    end
end

puts two_sum([1,2,3,4,9, 20, 11, 8], 10)
