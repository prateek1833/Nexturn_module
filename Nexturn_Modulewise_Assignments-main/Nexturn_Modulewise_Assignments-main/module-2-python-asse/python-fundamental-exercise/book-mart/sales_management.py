
from models import books, sales, Transaction

def sell_book(customer_name, book_title, quantity):
    try:
        quantity = int(quantity)
        for book in books:
            if book.title.lower() == book_title.lower():
                if book.quantity >= quantity:
                    book.quantity -= quantity
                    sales.append(Transaction(customer_name, book_title, quantity))
                    return f"Sale successful! Remaining quantity: {book.quantity}"
                else:
                    return f"Error: Only {book.quantity} copies available."
        return "Error: Book not found."
    except ValueError:
        return "Invalid input! Quantity must be a number."

def view_sales():
    if not sales:
        return "No sales records found."
    return "\n".join(sale.display() for sale in sales)