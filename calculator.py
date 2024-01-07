def calculate(expression):
    try:
        a, op, b = expression.split()
        a = int(a)
        b = int(b)
        if a < 1 or a > 10 or b < 1 or b > 10:
            return "Ошибка: числа должны быть от 1 до 10 включительно"
        if op == '+':
            return str(a + b)
        elif op == '-':
            return str(a - b)
        elif op == '*':
            return str(a * b)
        elif op == '/':
            return str(a // b)
        else:
            return "Ошибка: Некорректная арифметическая операция"
    except ValueError:
        return "Ошибка: Некорректный формат выражения"

def main(input):
    if "Ошибка" in input:
        return input
    return calculate(input)

# Пример использования
while True:
    expression = input()
    result = main(expression)
    if "Ошибка" in result:
        print(result)
        break
    print(result)