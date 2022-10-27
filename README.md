# AllGoRithms

AllGoRithms is a Golang project where I am putting into practice all my Algorithm studies.

It seems like a lot of people are enjoying this project as I have at least 100 clones a day since I started adding new algorithms.

So, I decided to add some information about the project.

## Structure

The project has the following tree:


```
all-go-rithms
├─ .vscode
│  └─ launch.json
├─ go.mod
├─ model
│  ├─ collection.go
│  └─ heap.go
├─ sort
│  ├─ counting_sort.go
│  ├─ heap_sort.go
│  ├─ insertion_sort.go
│  ├─ merge_sort.go
│  └─ quicksort.go
│  └─ ....
└─ test
   └─ sort
      ├─ counting_sort_test.go
      ├─ heapsort_test.go
      └─ quicksort_test.go
      └─ ....

```


As you can see, by now the project has only sorting algorithms. In the same way that my studies go to other Algorithms, we will have more directories.


## Tests

To run the tests, I recommend to use the following command:

```bash
go test $(find -name "*test.go*")
```

Since all my tests end with `test.go`, it will run all of them.

## Study references

[Introduction To Algorithms](https://www.amazon.com.br/Introduction-Algorithms-Thomas-H-Cormen/dp/0262033844/ref=asc_df_0262033844/?tag=googleshopp00-20&linkCode=df0&hvadid=379787788238&hvpos=&hvnetw=g&hvrand=313458358507001429&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=1001552&hvtargid=pla-422923046610&psc=1)
