import os
import json
import logging
import subprocess
from typing import Dict

# Define logging configuration
logging.basicConfig(level=logging.INFO, format='%(asctime)s [%(levelname)s] %(message)s')

# Define constants for Terraform configuration files
TERRAFORM_CONFIG = 'terraform.tf'
TERRAFORM_VAR = 'terraform.tfvars'
TERRAFORM_BACKEND = 'terraform.tfbackend'

def load_config(file_path: str) -> Dict:
    """
    Load configuration from a JSON file.
    
    Args:
    file_path (str): Path to the configuration file.
    
    Returns:
    Dict: Loaded configuration as a dictionary.
    """
    try:
        with open(file_path, 'r') as f:
            config = json.load(f)
        return config
    except FileNotFoundError:
        logging.error(f"Configuration file not found: {file_path}")
        return None
    except json.JSONDecodeError as e:
        logging.error(f"Failed to parse configuration file: {e}")
        return None

def run_terraform(command: str) -> str:
    """
    Run Terraform command and return the output.
    
    Args:
    command (str): Terraform command to run.
    
    Returns:
    str: Output of the Terraform command.
    """
    try:
        result = subprocess.run(['terraform', command], check=True, capture_output=True)
        return result.stdout.decode('utf-8')
    except subprocess.CalledProcessError as e:
        logging.error(f"Failed to run Terraform command: {e}")
        return None