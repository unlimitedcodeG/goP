import inspect


def get_caller_module():
    stack = inspect.stack()
    for frame in stack:
        module = inspect.getmodule(frame[0])
        if module and module.__name__.startswith("maa"):
            return module.__name__
    return None


caller = get_caller_module()

if caller == "maa":
    print("maa")
elif caller and caller.startswith("maa."):
    print(f"Imported via `{caller}`")
else:
    print("Unknown import source")
