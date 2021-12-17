from argon2 import PasswordHasher, exceptions
import datetime
import multiprocessing


def main():
    start = datetime.datetime.now()
    pool = multiprocessing.Pool()
    pool.map(cracker, range(0, 100000))
    pool.close()

    diff = datetime.datetime.now() - start
    print(f"Done in {diff.microseconds}")


def cracker(i: int):
    hash = "$argon2id$v=19$m=16,t=2,p=1$MTExMTExMTE$SzwiO6Uix4CqutzHAncBwQ"
    ph = PasswordHasher()
    try:
        ph.verify(hash, f"password{i}")
    except exceptions.VerifyMismatchError:
        print(".")


if __name__ == "__main__":
    main()
