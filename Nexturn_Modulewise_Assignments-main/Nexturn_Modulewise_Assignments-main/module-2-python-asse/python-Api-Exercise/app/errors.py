from flask import Blueprint, jsonify

errors_bp = Blueprint('errors', __name__)

@errors_bp.app_errorhandler(500)
def internal_error(error):
    return jsonify(error="Internal Server Error", message=str(error)), 500

@errors_bp.app_errorhandler(404)
def not_found_error(error):
    return jsonify(error="Resource Not Found", message="The requested resource does not exist"), 404