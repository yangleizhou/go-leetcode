## 素数

### 定义
- 如果一个数如果只能被 1 和它本身整除，那么这个数就是素数

### 题目
- 返回区间 [2, n) 中有几个素数

### 方法
#### 方法1：遍历i->n判断n是不是素数
- 时间复杂度: `O(n^2)`

#### 方法2：遍历i->sqrt(n)是不是素数
- 根据因子的对称性
- 假设 `n = 12` 只需要遍历
    12 = 2 × 6    
    12 = 3 × 4    
    12 = sqrt(12) × sqrt(12)    
    12 = 4 × 3    
    12 = 6 × 2    
- 外层的for 循环也可以根据因子的对称性 遍历[2,sqrt(n))
- 时间复杂度：`O(n) = O(sqrt(n)*sqrt(n))`

#### 方法3：素数的倍数下一定是合数
- 2 素数 2*2=4 2*3=6 2*4=8 2*5=10 ... 都是合数  
  3 素数 3*2=6 3*3=9 3*4=12... 都是合数
  5 素数 5*2=10 5*3=15 5*4=20... 都是合数
  ......
- 时间复杂度：O(n^2 * loglogn)  
    - n * (n/2+n/3+n/4+n/5+n/6+n/7+....)    
    = n^2*loglogn

#### 方法4：埃拉托色尼筛选法（埃氏筛法 sieve of Eratosthenes）
- 根据因子对称性 2*3=3*2=6 2*5=5*2=10 外层遍历可以 [2,sqrt(n))

- 时间复杂度：O(n * loglogn)  
    - n * (n/2+n/3+n/4+n/5+n/6+n/7+....)    
    = n*loglogn





#### 测试结果（n = 1000000）
    - 方法1：
    prime one =  9592
    --- PASS: TestCountPrimeOne (4.05s)

    - 方法2：
    prime two =  9592
    --- PASS: TestCountPrimeTwo (0.07s)

    prime two =  78498 (n=10000000)
    --- PASS: TestCountPrimeTwo (0.66s)

    - 方法3：
    prime three =  9592
    --- PASS: TestCountPrimeThree (0.02s)

    prime three =  78498 (n=10000000)
    --- PASS: TestCountPrimeThree (0.42s)

    - 方法4：
    prime four =  9592
    --- PASS: TestCountPrimeFour (0.02s)

    prime three =  78498
    --- PASS: TestCountPrimeFour (0.37s)