from flask import Flask
from flask_sqlalchemy import SQLAlchemy # type: ignore

db = SQLAlchemy()

def create_app():
    app = Flask(__name__)
    app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///books.db'
    app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

    db.init_app(app)

    with app.app_context():
        from . import routes, errors
        db.create_all()

        # Register blueprints
        app.register_blueprint(routes.bp)
        app.register_blueprint(errors.errors_bp)

    return app