import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import ProfileList from './ProfileList';
import SearchUser from './SearchUser';
import EditUser from './EditUser';
import Suggestion from './Suggestion';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: 80
  },
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
}));

export default function Profile() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={3}>
        <Grid item xs={6}>
          <Paper className={classes.paper}>
              <ProfileList></ProfileList>
          </Paper>
        </Grid>
        <Grid item xs={6}>
          <Paper className={classes.paper}>
              <SearchUser></SearchUser>
          </Paper>
        </Grid>
        <Grid item xs={6}>
          <Paper className={classes.paper}>
              <EditUser></EditUser>
          </Paper>
        </Grid>
        <Grid item xs={6}>
          <Paper className={classes.paper}>
              <Suggestion></Suggestion>
          </Paper>
        </Grid>
        
      </Grid>
    </div>
  );
}
