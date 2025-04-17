// Why can't we go with a typical buble sort 
// The most standard one


package bubblesort

func bubbleSort(arr []int){
  var n int = len(arr)

  for i:=0; i<n; i++{
    for j:=0; j<n-i-1; j++{
      if arr[j] > arr[j+1]{
        arr[j], arr[j+1] = arr[j+1], arr[j]
      }
    }
  }
}
