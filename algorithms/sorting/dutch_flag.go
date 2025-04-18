package sorting

func DF(nums []int){
    low := 0
    high := len(nums) - 1
    mid := 0

    for mid <= high{
        if nums[mid] == 0{
            nums[low], nums[mid] = nums[mid], nums[low]
            low++
            mid++
        }else if nums[mid] == 1{
            mid++
        }else{
            nums[mid], nums[high] = nums[high], nums[mid]
            high--
        }
    }
}
