function twoSum(nums: number[], target: number): number[] {
  const hasher: Record<number, number> = {};

  for (let index = 0; index < nums.length; index++) {
    if (nums[index] in hasher) {
      return [hasher[nums[index]], index];
    }
    hasher[target - nums[index]] = index;
  }

  return [];
}

console.log(twoSum([3, 12, 4, 7], 7));
