import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import InputLabel from '@material-ui/core/InputLabel';
import Input from '@material-ui/core/Input';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import AddIcon from '@material-ui/icons/Add';
import axios from 'axios';


const styles = theme => ({
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
});

class DialogSelect extends Component{
  state = {open: false, name: '', lists:[]}

   handleChange = (event) => {
    // setAge(Number(event.target.value) || '');
    this.setState({name: event.target.value || ''})
  };

   handleClickOpen = () => {
    // setOpen(true);
    this.setState({open:true})
    axios.get('/user/' + localStorage.getItem('username')).then(
        res => {
            this.setState({lists: res.data.lists});
            console.log(this.state.lists);
        }
    ).catch(
        err => {
            console.log(err);
        }
    )
  };

   handleClose = () => {
    // setOpen(false);
    this.setState({open:false})
  };

  handleAddItem = () => {
      this.setState({open:false});
      const data = {name:this.props.name, link:this.props.link, price:this.props.price, interest:6}
      var i;
      for (i = 0; i < this.state.lists.length; i++) {
        if (this.state.lists[i].n == this.state.name){
            axios.post('/user/list/' + this.state.lists[i]._id + '/gift', data).then(
                res => {
                    console.log('success');
                }
            ).catch(
                err => {
                    console.log('failure');
                }
            )
        }
      }
      
  }

  render(){
    const {classes} = this.props;
    return (
        <div>
          <AddIcon color="primary" onClick={this.handleClickOpen}></AddIcon>
          <Dialog disableBackdropClick disableEscapeKeyDown open={this.state.open} onClose={this.handleClose}>
            <DialogTitle>لیست مورد نظر را انتخاب کنید</DialogTitle>
            <DialogContent>
              <form className={classes.container}>
                <FormControl className={classes.formControl}>
                  <InputLabel id="list-select">نام لیست</InputLabel>
                  <Select
                    labelId="demo-dialog-select-label"
                    id="list-selectt"
                    value={this.state.name}
                    onChange={this.handleChange}
                    // input={<Input />}
                  >
                    {/* <MenuItem value="">
                      <em> </em>
                    </MenuItem>
                    <MenuItem value={10}>Ten</MenuItem>
                    <MenuItem value={20}>Twenty</MenuItem>
                    <MenuItem value={30}>Thirty</MenuItem> */}
                    {
                        this.state.lists.map((list) => (
                            <MenuItem value={list.n}>{list.n}</MenuItem>
                        ))
                    }
                  </Select>
                </FormControl>
              </form>
            </DialogContent>
            <DialogActions>
              <Button onClick={this.handleClose} color="secondary">
                بیخیال!
              </Button>
              <Button onClick={this.handleAddItem} color="primary">
                تایید
              </Button>
            </DialogActions>
          </Dialog>
        </div>
      );
  }
}

export default withStyles(styles)(DialogSelect)