from app import create_app, db
from app.models import Book

app = create_app()

with app.app_context():
    db.drop_all()
    db.create_all()

    sample_books = [
        {"title": "The Great Gatsby", "author": "F. Scott Fitzgerald", "published_year": 1925, "genre": "Fiction"},
        {"title": "To Kill a Mockingbird", "author": "Harper Lee", "published_year": 1960, "genre": "Fiction"},
        {"title": "Dune", "author": "Frank Herbert", "published_year": 1965, "genre": "Sci-Fi"},
    ]

    for book in sample_books:
        db.session.add(Book(**book))
    db.session.commit()