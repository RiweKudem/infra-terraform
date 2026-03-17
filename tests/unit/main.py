import os
import click
import yaml
from terralink import Terralink

@click.group()
def main():
    """Main entry point for the infra-terraform CLI."""
    pass

@main.command()
@click.option('--config', '-c', help='Path to the configuration file.')
@click.option('--workspace', '-w', help='Name of the workspace to use.')
def init(config, workspace):
    """Initialize the Terraform workspace."""
    config_file = os.path.abspath(config)
    with open(config_file, 'r') as f:
        config_data = yaml.safe_load(f)

    if config_data['workspace'] != workspace:
        click.echo('Workspace mismatch: {} vs {}'.format(config_data['workspace'], workspace))

    tl = Terralink(config_file)
    tl.init_workspace()

@main.command()
@click.option('--config', '-c', help='Path to the configuration file.')
@click.option('--workspace', '-w', help='Name of the workspace to use.')
@click.option('--apply', '-a', help='Apply the Terraform configuration.', is_flag=True)
def plan(config, workspace, apply):
    """Get the Terraform plan."""
    config_file = os.path.abspath(config)
    with open(config_file, 'r') as f:
        config_data = yaml.safe_load(f)

    if config_data['workspace'] != workspace:
        click.echo('Workspace mismatch: {} vs {}'.format(config_data['workspace'], workspace))

    tl = Terralink(config_file)
    if apply:
        tl.apply_plan()
    else:
        tl.get_plan()

if __name__ == '__main__':
    main()