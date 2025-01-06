from flask import Blueprint, request, jsonify
from .models import Book, db
from .validation import validate_book_data

bp = Blueprint('routes', __name__)

@bp.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()
    error = validate_book_data(data)
    if error:
        return jsonify(error=error), 400

    new_book = Book(**data)
    db.session.add(new_book)
    db.session.commit()
    return jsonify(message="Book added successfully", book_id=new_book.id), 201

@bp.route('/books', methods=['GET'])
def get_books():
    books = Book.query.all()
    return jsonify([book.to_dict() for book in books])

@bp.route('/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify(error="Book not found", message="No book exists with the provided ID"), 404
    return jsonify(book.to_dict())

@bp.route('/books/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify(error="Book not found", message="No book exists with the provided ID"), 404

    data = request.get_json()
    error = validate_book_data(data)
    if error:
        return jsonify(error=error), 400

    for key, value in data.items():
        setattr(book, key, value)
    db.session.commit()
    return jsonify(message="Book updated successfully")

@bp.route('/books/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    book = Book.query.get(book_id)
    if not book:
        return jsonify(error="Book not found", message="No book exists with the provided ID"), 404

    db.session.delete(book)
    db.session.commit()
    return jsonify(message="Book deleted successfully")