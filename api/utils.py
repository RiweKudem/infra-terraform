import logging
import sys
from pathlib import Path
import yaml

def load_config(config_file):
    try:
        with open(config_file, 'r') as file:
            config = yaml.safe_load(file)
            return config
    except yaml.YAMLError as e:
        logging.error(f"Error loading config file: {e}")
        sys.exit(1)

def get_parent_dir(n=1):
    current_dir = Path(__file__).resolve().parent
    for _ in range(n):
        current_dir = current_dir.parent
    return current_dir

def get_project_root():
    current_dir = Path(__file__).resolve().parent
    while not (current_dir / 'infra-terraform').exists():
        current_dir = current_dir.parent
    return current_dir

def configure_logging():
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s [%(levelname)s] %(message)s',
        handlers=[
            logging.FileHandler('log.txt'),
            logging.StreamHandler()
        ]
    )