from book_management  import add_book, view_books, search_book
from customer_management import add_customer, view_customers
from sales_management import sell_book, view_sales

def main():
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        
        choice = input("Enter your choice: ")
        
        if choice == '1':
            print("\n1. Add Book\n2. View Books\n3. Search Book")
            book_choice = input("Enter your choice: ")
            
            if book_choice == '1':
                title = input("Title: ")
                author = input("Author: ")
                price = input("Price: ")
                quantity = input("Quantity: ")
                book_management.add_book(title, author, price, quantity)
            elif book_choice == '2':
                book_management.view_books()
            elif book_choice == '3':
                search_term = input("Enter title or author to search: ")
                book_management.search_book(search_term)
        
        elif choice == '2':
            print("\n1. Add Customer\n2. View Customers")
            customer_choice = input("Enter your choice: ")
            
            if customer_choice == '1':
                name = input("Name: ")
                email = input("Email: ")
                phone = input("Phone: ")
                customer_management.add_customer(name, email, phone)
            elif customer_choice == '2':
                customer_management.view_customers()
        
        elif choice == '3':
            print("\n1. Sell Book\n2. View Sales Records")
            sales_choice = input("Enter your choice: ")
            
            if sales_choice == '1':
                customer_name = input("Customer Name: ")
                book_title = input("Book Title: ")
                quantity = input("Quantity: ")
                sales_management.sell_book(customer_name, book_title, quantity)
            elif sales_choice == '2':
                sales_management.view_sales()
        
        elif choice == '4':
            print("Exiting the system.")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main()
