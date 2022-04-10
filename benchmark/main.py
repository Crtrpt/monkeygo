import time

def Fibonacci(n):
    if n < 0:
        print("Incorrect input")
    elif n == 0:
        return 0
    elif n == 1 or n == 2:
        return 1
    else:
        return Fibonacci(n-1) + Fibonacci(n-2)
 

start= time.time()
res=Fibonacci(35)
end= time.time()
print("engine={}, result={}, duration={}".format("python",res,end-start))