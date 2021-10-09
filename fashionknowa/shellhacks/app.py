import string
import re
import json
import os
from flask import Flask, flash, render_template, redirect, request, session, make_response, session, redirect, url_for, send_from_directory
import requests
import os
import numpy as np
import pandas as pd
import pytesseract
import cv2
from werkzeug.utils import secure_filename
#from autoscraper import AutoScraper

app = Flask(__name__)
app.config['UPLOAD_FOLDER'] = 'static/uploads/'
app.config['MAX_CONTENT_LENGTH'] = 16 * 1024 * 1024
app.secret_key = os.environ.get('APP_SECRET_KEY')

ALLOWED_EXTENSIONS = set(['png', 'jpg', 'jpeg', 'gif'])

sustainable = ['cotton', 'linen', 'silk', 'wool', 'bamboo', 'hemp',
               'lyocell', 'nylon', 'rubber', 'cashmere', 'cupro', 'denim']

non_sustainable = ['leather', 'rayon', 'viscose', 'polyester',
                   'elastane', 'elasthane', 'spandex', 'chiffon', 'acrylic', 'georgette']


def allowed_file(filename):
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS


@app.route('/')
def upload_form():
    return render_template('upload.html')


@app.route('/about')
def about():
    return render_template('about.html')


@app.route('/sustainability')
def why():
    return render_template('why.html')


@app.route('/', methods=['POST'])
def upload_image():
    website = request.form['text']
    '''if website:
        scraper = AutoScraper()
        sus = scraper.build(website, sustainable)
        nsus = scraper.build(website, non_sustainable)
        print(sus, nsus)
    else:
        print(No Info)'''
    if 'file' not in request.files:
        return render_template('upload.html', flash="No file part")
    file = request.files['file']
    if file.filename == '':
        return render_template('upload.html', flash='No image selected for uploading')
    if file and allowed_file(file.filename):
        filename = secure_filename(file.filename)
        path = os.path.join(app.config['UPLOAD_FOLDER'], filename)
        file.save(path)
        image = cv2.imread(path)
        image = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)
        pytesseract.pytesseract.tesseract_cmd = '/app/.apt/usr/bin/tesseract'
        text = pytesseract.image_to_string(image)
        os.remove(path)
        lst = list(text.lower().strip().split())
        s = []
        ns = []
        for l in lst:
            if l in sustainable:
                s.append(l)
            elif l in non_sustainable:
                ns.append(l)
        ty = -2
        if len(s) > len(ns):
            ty = 1
        elif len(s) < len(ns):
            ty = -1
        elif len(s) == len(ns) and len(s) > 0:
            ty = 0
        return render_template('analysis.html', s=s, ns=ns, type=ty)
    else:
        return render_template('upload.html', flash='Allowed image types are png, jpg, jpeg, gif')


if __name__ == '__main__':
    app.run(debug=True, port=5500)
