# -*- coding: utf-8 -*-
#
# Configuration file for the Sphinx documentation builder.
#
# This file does only contain a selection of the most common options. For a
# full list see the documentation:
# http://www.sphinx-doc.org/en/stable/config

# -- Path setup --------------------------------------------------------------

# If extensions (or modules to document with autodoc) are in another directory,
# add these directories to sys.path here. If the directory is relative to the
# documentation root, use os.path.abspath to make it absolute, like shown here.
#
import os
import re


# -- Project information -----------------------------------------------------

project = "Flytectl"
copyright = "2021, Flyte"
author = "Flyte"

# The full version, including alpha/beta/rc tags
release = re.sub("^v", "", os.popen("git describe").read().strip())
version = release


# -- General configuration ---------------------------------------------------

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
extensions = [
    "sphinx.ext.autosummary",
    "sphinx.ext.autosectionlabel",
    "sphinx.ext.intersphinx",
    "sphinx.ext.todo",
    "sphinx.ext.viewcode",
    "sphinx.ext.doctest",
    "sphinx.ext.coverage",
    "sphinx-prompt",
    "sphinx_copybutton",
    "sphinx_fontawesome",
    "sphinxcontrib.youtube",
    "sphinx_reredirects",
    "sphinx_panels",
]

# build the templated autosummary files
autosummary_generate = True

# autosectionlabel throws warnings if section names are duplicated.
# The following tells autosectionlabel to not throw a warning for
# duplicated section names that are in different documents.
autosectionlabel_prefix_document = True

# Add any paths that contain templates here, relative to this directory.
templates_path = ["_templates"]

# The suffix(es) of source filenames.
# You can specify multiple suffix as a list of string:
#
# source_suffix = ['.rst', '.md']
source_suffix = ".rst"

# The master toctree document.
master_doc = "index"

# The language for content autogenerated by Sphinx. Refer to documentation
# for a list of supported languages.
#
# This is also used if you do content translation via gettext catalogs.
# Usually you set "language" from the command line for these cases.
language = None

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This pattern also affects html_static_path and html_extra_path .
exclude_patterns = [
    u"_build",
    "Thumbs.db",
    ".DS_Store",
    "docs_index.rst",
    "overview.rst",
]

# The name of the Pygments (syntax highlighting) style to use.
pygments_style = "tango"
pygments_dark_style = "native"

# -- Options for HTML output -------------------------------------------------

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_theme = "furo"
html_title = "Flyte"
html_logo = "flyte_circle_gradient_1_4x4.png"
html_favicon = "flyte_circle_gradient_1_4x4.png"

announcement = """
📢 This is the old documentation for Flyte.
Please visit the new documentation <a href="https://docs.flyte.org">here</a>.
"""

html_theme_options = {
    "light_css_variables": {
        "color-brand-primary": "#4300c9",
        "color-brand-content": "#4300c9",
        "color-announcement-background": "#FEE7B8",
        "color-announcement-text": "#535353",
    },
    "dark_css_variables": {
        "color-brand-primary": "#9D68E4",
        "color-brand-content": "#9D68E4",
        "color-announcement-background": "#493100",
    },
    "announcement": announcement,
}

html_context = {
    "home_page": "https://docs.flyte.org",
    # custom flyteorg furo theme options
    "github_repo": "flytectl",
    "github_username": "flyteorg",
    "github_commit": "master",
    "docs_path": "docs/source",  # path to documentation source
}

# The default sidebars (for documents that don't match any pattern) are
# defined by theme itself.  Builtin themes are using these templates by
# default: ``['localtoc.html', 'relations.html', 'sourcelink.html',
# 'searchbox.html']``.
# html_sidebars = {"**": ["logo-text.html", "globaltoc.html", "localtoc.html", "searchbox.html"]}


# Add any paths that contain custom static files (such as style sheets) here,
# relative to this directory. They are copied after the builtin static files,
# so a file named "default.css" will overwrite the builtin "default.css".
html_static_path = ["_static"]
html_css_files = ["custom.css"]

# Custom sidebar templates, must be a dictionary that maps document names
# to template names.
#


# -- Options for HTMLHelp output ---------------------------------------------

# Output file base name for HTML help builder.
htmlhelp_basename = "flytectldoc"


# -- Options for LaTeX output ------------------------------------------------

latex_elements = {
    # The paper size ('letterpaper' or 'a4paper').
    #
    # 'papersize': 'letterpaper',
    # The font size ('10pt', '11pt' or '12pt').
    #
    # 'pointsize': '10pt',
    # Additional stuff for the LaTeX preamble.
    #
    # 'preamble': '',
    # Latex figure (float) alignment
    #
    # 'figure_align': 'htbp',
}

# Grouping the document tree into LaTeX files. List of tuples
# (source start file, target name, title,
#  author, documentclass [howto, manual, or own class]).
latex_documents = [
    (master_doc, "flytectl.tex", "flytectl Documentation", "Flyte", "manual"),
]


# -- Options for manual page output ------------------------------------------

# One entry per manual page. List of tuples
# (source start file, name, description, authors, manual section).
man_pages = [(master_doc, "flytectl", "flytectl Documentation", [author], 1)]


# -- Options for Texinfo output ----------------------------------------------

# Grouping the document tree into Texinfo files. List of tuples
# (source start file, target name, title, author,
#  dir menu entry, description, category)
texinfo_documents = [
    (
        master_doc,
        "flytectl",
        "flytectl Documentation",
        author,
        "flytectl",
        "The one CLI for flyte.",
        "Miscellaneous",
    ),
]

# -- Options for intersphinx -------------------------------------------------
# intersphinx configuration
intersphinx_mapping = {
    "flyteidl": ("https://docs.flyte.org/en/latest/reference_flyteidl.html", None),
    "flyte": ("https://docs.flyte.org/en/latest", None),
}

if int(os.environ.get("ENABLE_SPHINX_REDIRECTS", 0)):
    # Redirects to the new docs site
    redirects = {
        "verbs.html": "https://docs.flyte.org/en/latest/flytectl/verbs.html",
        "nouns.html": "https://docs.flyte.org/en/latest/flytectl/nouns.html",
        "gen/*": "https://docs.flyte.org/en/latest/flytectl/$source.html",
        "contribute.html": "https://docs.flyte.org/en/latest/flytectl/contribute.html",
    }
