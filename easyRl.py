import time


class Test:
    def __init__(self):
        self.value = 42

    def hello(self):
        return "hello world"


obj = Test()


# 获取属性值
print(getattr(obj, "value"))  # 42


# 设置属性值

setattr(obj, "value", 100)

print(obj.value)

print(hasattr(obj, "value"))

method = getattr(obj, "hello")


def timer(func):
    def wrapper(*args, **kargs):
        start = time.time()
        result = func(*args, **kargs)
        end = time.time()
        print(f"{func.__name__} 执行时间:{end - start:.4f}s")

    return wrapper


@timer
def slow_function():
    time.sleep(1)
    print("func exec")


slow_function()
